'use strict'
var pricesApp = angular.module('pricesApp', ['ngResource']);

pricesApp.config(['$httpProvider', function($httpProvider) {
    $httpProvider.defaults.useXDomain = true;
    $httpProvider.defaults.headers.put['Content-Type'] = 'application/xml';
}]);
pricesApp.controller('pricesCtrl', ['$scope', '$http', '$window',
    function($scope, $http, $window) {

        $scope.$window = $window;

        $scope.prices = [
            {"tick": "RIL", "price": 101},
            {"tick": "TCS", "price": 121},
            {"tick": "INFY", "price": 131},
            
        ];
        
        $scope.selectedPrice = {
            "tick" : "",
            "price" : 0
        };
        $scope.onSelect = function(price){
            $scope.selectedPrice = price;
            $scope.selectedPrice.price = parseFloat(price.price);
            
        };
         $scope.onPriceChange = function(){
            console.log("Changed Price: ", $scope.selectedPrice);
            $http.put("http://localhost:9000/prices/"+$scope.selectedPrice.tick, $scope.selectedPrice).
              success(function(data, status, headers, config) {
                    console.log("Success");
					
                }).
                error(function(data, status, headers, config) {
                    $scope.error = true;
                    console.log("Error");
                });
        };
}]);