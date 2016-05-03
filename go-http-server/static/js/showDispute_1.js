/*******************************************************************************
 * Copyright (c) 2012, 2015 SunGard CSA LLC and others.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors:
 * SunGard CSA LLC - initial API and implementation and/or initial documentation
 *******************************************************************************/
'use strict'
var ippUIMashup = angular.module('ippUIMashup', ['ngResource']);

ippUIMashup.config(['$httpProvider', function($httpProvider) {
    $httpProvider.defaults.useXDomain = true;
    $httpProvider.defaults.headers.put['Content-Type'] = 'application/xml';
}]);
ippUIMashup.controller('disputeCtrl', ['$scope', '$http', '$window',
    function($scope, $http, $window) {

        $scope.$window = $window;

        $scope.Dispute = {};
        $scope.MemberData = {};
        var x2js = new X2JS();
        $scope.urlParam = function(name) {
            var results = new RegExp('[\\?&]' + name + '=([^&#]*)').exec(window.location.href);
            if (!results) {
                return 0;
            }
            return results[1] || 0;
        }

        $scope.saveDispute = function() {
                var disputeData = "<?xml version=\"1.0\" encoding=\"UTF-8\"?><Dispute>" + x2js.json2xml_str($scope.Dispute) + "</Dispute>";
                console.log("Dispute to Submit: " + disputeData);
                /*This section of scripts write data back to IPP from $scope disputeData*/
                var res1 = $scope.callbackURL + '/outData/Dispute';
                $http.put(res1, disputeData).
                success(function(data, status, headers, config) {
                    console.log("Success");
					$scope.sleep(3000);
					IppProcessPortalClient.completeActivity();
                }).
                error(function(data, status, headers, config) {
                    $scope.error = true;
                    console.log("Error");
                });
			}	
            
            $scope.suspendDispute = function() {
                var disputeData = "<?xml version=\"1.0\" encoding=\"UTF-8\"?><Dispute>" + x2js.json2xml_str($scope.Dispute) + "</Dispute>";
                console.log("Dispute to Submit: " + disputeData);
                /*This section of scripts write data back to IPP from $scope disputeData*/
                var res1 = $scope.callbackURL + '/outData/Dispute';
                $http.put(res1, disputeData).
                success(function(data, status, headers, config) {
                    console.log("Success");
					$scope.sleep(3000);
					IppProcessPortalClient.suspendActivity(true);
                }).
                error(function(data, status, headers, config) {
                    $scope.error = true;
                    console.log("Error");
                });
			}
			/**
             * Delay for a number of milliseconds
             */
        $scope.sleep = function(delay) {
            var start = new Date().getTime();
            while (new Date().getTime() < start + delay);
        }


        $scope.sleep(1000);

        $scope.callbackURL = $scope.urlParam("ippInteractionUri");
        //To ensure that $resource maintains port no.
        if ($scope.callbackURL) {
            $scope.callbackURL = $scope.callbackURL.replace("\\", "");
        }
        console.log($scope.callbackURL);

        /*This section of scripts fetches data from IPP and set it to $scope data*/
        var res2 = $scope.callbackURL + '/inData/Dispute55';
        $http.get(res2).
        success(function(data, status, headers, config) {
            console.log("Data is:" + data);
            var json = x2js.xml_str2json(data);//{Dispute: {disputeID: 111, description: "Dispute Description"}}; 
            //
            console.log("JSON Data is:" + json.Dispute);
            $scope.Dispute.disputeID = json.Dispute.disputeID;
            $scope.Dispute.description = json.Dispute.description;
        }).
        error(function(data, status, headers, config) {
            $scope.error = true;
        });
    }
]);