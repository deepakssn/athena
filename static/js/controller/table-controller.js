angular.module('app', [])
  .controller('tableController', ['$scope', '$http', function($scope, $http) {
    console.log('app works');
    $scope.availableFields = [];
    $scope.searchKeyword = '';
    $scope.sortField = '';
    // $http.get("https://api.myjson.com/bins/73fkv")
    $http.get("/report")
      .then(function(response) {
        $scope.response = response.data.gameDetails;
        var keys = [];
        for (var i = 0; i < $scope.response.length; i++) {
          Object.keys($scope.response[i]).forEach(function(key) {
            if ($scope.availableFields.indexOf(key) == -1) {
              $scope.availableFields.push(key);
            }
          });
        }
        $scope.sortField = $scope.availableFields[0];
      });
    $scope.selectSortField = function(sortField) {
      $scope.sortField = sortField;
    }
  }]);