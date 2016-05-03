/**
 * 
 */
var disputeIppApp = angular.module('disputeIppApp', ['IPPUiMashupUtils']);

disputeIppApp.controller('disputeController', ['$scope', 'IPPUiMashupService',
  function($scope,IPPUiMashupService) {
    $scope.Dispute = {};
    $scope.MemberData = {};
    $scope.saveDispute = function() {

      IPPUiMashupService.setOutputData('Dispute',$scope.Dispute, 'Dispute').then(
        function(data) {
          //success
          IPPUiMashupService.completeActivity();
        },
        function(data){
          //error
          $scope.error = data;
        }
      );
    };

    $scope.getInputData = function() {
      IPPUiMashupService.getInputData('Dispute').then(
        function(data) {
          //success
          if (angular.isDefined(data) && angular.isDefined(data.Dispute) && angular.isObject(data)) {
            angular.copy(data.Dispute, $scope.Dispute);
          }
        },
        function(data){
          //error
          $scope.error = data;
        }
      );
    }

    $scope.getInputData();

    $scope.suspendDispute = function() {
       IPPUiMashupService.setOutputData('Dispute',$scope.Dispute, 'Dispute').then(
        function(data) {
          //success
          IPPUiMashupService.suspendActivity();
        },
        function(data){
          //error
          $scope.error = data;
        }
      );

    }
  }
]);
