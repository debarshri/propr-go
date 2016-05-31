// Define the `phonecatApp` module
var appProp = angular.module('propr-app', ['ngRoute']);

appProp.config(['$locationProvider', '$routeProvider',
            function config($locationProvider, $routeProvider) {
              $locationProvider.hashPrefix('!');

              $routeProvider

              .when('/', {
                  template: '<div ng-controller="ExampleController" class="form-group"><form novalidate> '+
                                                '<legend>App Name</legend>'+
                                                '<input type="text" ng-model="user.property" class="form-control" placeholder="property name" />'+
                                                '<br /><input type="button" ng-click="search(user)" value="Search" class="btn btn-default" /> '+
                                                '</form><pre ng-show="result.property != null" ng-model="result">{{result.property}}</pre></div></div>'})

            .when('/create-property', {
                  template: '<div ng-controller="CreatePropertyController" class="form-group">'+
                                                  '<form novalidate><legend>App Name</legend>'+
                                                  '<input type="text" ng-model="property.property" class="form-control" placeholder="property name" />'+
                                                  '<br /><input type="button" ng-click="create(property)" value="Create Property" class="btn btn-default" />'+
                                                  '</form>'+
                                                  '<pre ng-show="result.property != null"  ng-model="result">{{result.property}}</pre></div></div>' })

            .when('/put-value', {
                  template: '<div ng-controller="PutValueController" class="form-group">'+
                                                  '<form novalidate><label>App Name</label>'+
                                                  '<input type="text" ng-model="property.name" class="form-control" placeholder="property name" />'+
                                                  '<br /><input type="text" ng-model="property.key" class="form-control" placeholder="property key" />'+
                                                  '<br /><input type="text" ng-model="property.value" class="form-control" placeholder="property value" />'+
                                                  '<br /><input type="button" ng-click="create(property)" value="Save" class="btn btn-default" />'+
                                                  '</form>'+
                                                  '<pre ng-show="result.property != null" ng-model="result">{{result.property}}</pre></div></div>'
                }).
                otherwise('/');
            }
          ]);

appProp.controller('ExampleController', function ExampleController($scope, $http) {

    function successCallback(response){
        $scope.result.property = response.data
    };

    function errorCallback(response){
        $scope.result.property =  response.data
    };

    $scope.result = {};

    $scope.search = function(user) {
        $http.get('/get/'+user.property).then(successCallback, errorCallback);
    };

});


appProp.controller('CreatePropertyController', function CreatePropertyController($scope, $http) {

    function successCallback(response){
        $scope.result.property =  response.data
    };

    function errorCallback(response){
        $scope.result.property =  response.data
    };

    $scope.result = {};

    $scope.create = function(property) {
        $http.post('/create/'+property.property).then(successCallback, errorCallback);
    };

});


appProp.controller('PutValueController', function CreatePropertyController($scope, $http) {

function successCallback(response){
    $scope.result.property =  response.data
};

function errorCallback(response){
    $scope.result.property =  response.data
};

$scope.result = {};

$scope.create = function(property) {
    $http.post('/add/'+property.name+'/'+property.key+'/'+property.value).then(successCallback, errorCallback);
};

});




