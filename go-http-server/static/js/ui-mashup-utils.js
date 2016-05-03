/***
 * Angular Module that has utility methods for IPP based Mashups
 */
var ippUiMashupUtils = angular.module("IPPUiMashupUtils", []);

ippUiMashupUtils.config(['$httpProvider', function($httpProvider) {
  $httpProvider.defaults.useXDomain = true;
  $httpProvider.defaults.headers.put['Content-Type'] = 'application/xml';
}]);

/***
 * Angular service that has utility methods for IPP based Mashups
 */
ippUiMashupUtils.factory('IPPUiMashupService', ['$window', '$document', '$q', '$http', 'XmlJsonService', function ($window, $document, $q, $http, XmlJsonService) {
  /**
   * Utility method to get URL parameter values from the window location.
   * @param parameterName a URI query parameter name
   * @returns {*} 0 if parameter is not available or the value of the parameter if available
     */
  function getUrlParameter(parameterName) {
    console.log($window.location);
    var results = new RegExp('[\\?&]' + parameterName + '=([^&#]*)').exec($window.location.href);
    if (!results) {
      return 0;
    }
    return results[1] || 0;
  }

  var IPP_INTERACTION_URI_PARAMETER = 'ippInteractionUri';
  /***
   * Gets the IPP interaction URI
   * @returns {*}
     */
  function getIppInteractionUri() {
    var ippInteractionUri = getUrlParameter(IPP_INTERACTION_URI_PARAMETER);
    if(!ippInteractionUri || ippInteractionUri === null) {
      return null;
    }
    ippInteractionUri = ippInteractionUri.replace("\\", "");
    if(!ippInteractionUri || ippInteractionUri === null) {
      return null;
    }
    return ippInteractionUri;
  }

  /**
   * Utility method to load the IPP Portal Client script asynchronously.
   * @returns {Element}
     */
  function loadIppClientScript() {
    var url = getUrlParameter('ippPortalBaseUri') + "/plugins/processportal/IppProcessPortalClient.js";
    if (url) {
      var script = document.querySelector("script[src*='" + url + "']");
      if (!script) {
        var heads = document.getElementsByTagName("head");
        if (heads && heads.length) {
          var head = heads[0];
          if (head) {
            script = document.createElement('script');
            script.setAttribute('src', url);
            script.setAttribute('type', "text/javascript");
            head.appendChild(script);
          }
        }
      }

      return script;
    }
  }

  function initialize() {
    loadIppClientScript();
  }

  /**
   * Builds XML String from JSON object and the root tag name
   * @type {string}
     */
  var XML_DECLARATION = '<?xml version="1.0" encoding="UTF-8"?>';
  function buildXmlData(jsonObject, xmlRootTagName) {
    if (!jsonObject || !xmlRootTagName) return null;
    var xmlStr = XmlJsonService.jsonToXml(jsonObject);
    if (!xmlStr || xmlStr === null) {
      xmlStr = '';
    }
    var overallXml = XML_DECLARATION + "<" + xmlRootTagName + ">" + xmlStr + "</" + xmlRootTagName + ">";
    return overallXml;
  }

  /***
   * Sets the output data of an IPP activity. Executes asynchronously and returns a promise.
   * @param outputDataAccessPointId
   * @param jsonObject
   * @param xmlRootTagName
   * @returns {*} a promise (asynchronous)
     */
  function setOutputData(outputDataAccessPointId, jsonObject, xmlRootTagName) {
    var deferred = $q.defer();

    if(!outputDataAccessPointId || outputDataAccessPointId === null) {
      deferred.reject('Output access point ID cannot be null or empty');
      return deferred.promise;
    }

    var xmlToSendToOutput = buildXmlData(jsonObject, xmlRootTagName);
    if(!xmlToSendToOutput || xmlToSendToOutput === null) {
      deferred.reject('XML to send to output access point cannot be null');
      return deferred.promise;
    }

    var ippInteractionUri = getIppInteractionUri();
    if(ippInteractionUri === null) {
      deferred.reject('IPP interaction URI cannot be null or empty');
      return deferred.promise;
    }

    var outputUri = ippInteractionUri + '/outData/' + outputDataAccessPointId;

    $http.put(outputUri, xmlToSendToOutput).
      success(function(data,status,headers,config){
        deferred.resolve(data);
    }).
      error(function(data, status, headers, config) {
        deferred.reject(data);
    });

    return deferred.promise;
  }

  /***
   * Gets input data to the activity. Executes asynchronously and returns a promise.
   * @param inputDataAccessPointId
   * @returns {*} a promise (asynchronous)
     */
  function getInputData(inputDataAccessPointId) {
    var deferred = $q.defer();

    if(!inputDataAccessPointId || inputDataAccessPointId === null) {
      deferred.reject('Input access point ID cannot be null or empty');
      return deferred.promise;
    }

    var ippInteractionUri = getIppInteractionUri();
    if(ippInteractionUri === null) {
      deferred.reject('IPP interaction URI cannot be null or empty');
      return deferred.promise;
    }

    var inputUri = ippInteractionUri + '/inData/' + inputDataAccessPointId;

    $http.get(inputUri).
    success(function(data,status,headers,config){
      var jsonData=null;
      if(angular.isDefined(data) && angular.isString(data))
      {
        jsonData = XmlJsonService.xmlToJson(data)
      }
      deferred.resolve(jsonData);
    }).
    error(function(data, status, headers, config) {
      deferred.reject(data);
    });

    return deferred.promise;
  }

  /***
   * Completes the activity
   */
  function completeActivity() {
    waitForIppClientToLoad().then(
      function() {
        $window.IppProcessPortalClient.completeActivity();
      }
    );
  }

/***
   * Suspends the activity
   */
  function suspendActivity() {
    waitForIppClientToLoad().then(
      function() {
        $window.IppProcessPortalClient.suspendActivity(true);
      }
    );
  }

  var waitCount = 0;
  var WAIT_MILLIS = 1000;
  var MAX_WAIT_COUNT = 10;
  var deferredIppClient =  $q.defer();
  /***
   * A method to check if IppProcessPortalClient is loaded.
   * If not, retry and wait.
   * @returns {*} a promise (asynchronous)
     */
  function waitForIppClientToLoad() {
    if(!angular.isDefined($window.IppProcessPortalClient)) {
      // Not loaded yet
      if(waitCount > MAX_WAIT_COUNT) {
        deferredIppClient.reject('Fatal Error: Unable to load IPP client after ' + MAX_WAIT_COUNT + ' attempts.' );
        return;
      }
      waitCount++;
      loadIppClientScript();
      setTimeout(waitForIppClientToLoad, WAIT_MILLIS);
      deferredIppClient.notify(waitCount);

    }
    else {
      // Loaded
      deferredIppClient.resolve();
    }
    return deferredIppClient.promise;
  }

  return {
    initialize: initialize,
    getUrlParameter: getUrlParameter,
    buildXmlData: buildXmlData,
    setOutputData: setOutputData,
    getInputData: getInputData,
    waitForIppClientToLoad: waitForIppClientToLoad,
    completeActivity:completeActivity,
    suspendActivity:suspendActivity
  };
}]);

/***
 * Angular service to handle XML, JSON conversions
 */
ippUiMashupUtils.factory("XmlJsonService", function () {
  var x2js = new X2JS();

  function jsonToXml(jsonObject) {
    return x2js.json2xml_str(jsonObject);
  }

  function xmlToJson(xmlString) {
    return x2js.xml_str2json(xmlString);
  }

  return {
    jsonToXml: jsonToXml,
    xmlToJson: xmlToJson
  };
});


