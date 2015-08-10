'use strict';

angular
  .module('demo', ['mwl.calendar', 'ui.bootstrap'])
  .controller('MainCtrl',  function ($scope, $modal, moment, $http) {

// var Reservations = $resource('/reservations');
// Reservations.get(function(eves){
//   var newone = eves.a1;
// })


// $http({ method: 'GET', url: 'http://localhost:3000/reservations' }).
//   success(function (data, status, headers, config) {
//     $scope.events.title = data.a1
//   }).
//   error(function (data, status, headers, config) {
//     // ...
//   });


    //These variables MUST be set as a minimum for the calendar to work
    $scope.calendarView = 'month';
    $scope.calendarDay = new Date();
    
    var res = $http.get('http://localhost:3000/reservation');
    res.success(function(data, status, headrs, config){
      if(data != null) {


        data.forEach(function(ele, index, array){
          // $scope.events[index].title = ele.title;
          // $scope.events[index].type = ele.type;
          // $scope.events[index].startsAt = ele.startsAt;
          // $scope.events[index].endsAt = ele.endsAt;
        
          $scope.events[index] = {
            title : ele.title,
            type : ele.type,
            startsAt : ele.startsAt,
            endsAt : ele.endsAt
          }



        })
        
      }


      

     // event.title = data[1].id;
   });


    // $scope.events = [
    //   {
    //     title: "newone",
    //     type: 'warning',
    //     startsAt: moment().startOf('week').subtract(2, 'days').add(8, 'hours').toDate(),
    //     endsAt: moment().startOf('week').add(1, 'week').add(9, 'hours').toDate()
    //   }, {
    //     title: "newone",
    //     type: 'info',
    //     startsAt: 12039485678000,
    //     endsAt: 1234576967000
    //   }, {
    //     title: "newone",
    //     type: 'important',
    //     startsAt: moment().startOf('day').add(7, 'hours').toDate(),
    //     endsAt: moment().startOf('day').add(19, 'hours').toDate(),
    //     recursOn: 'year'
    //   }
    // ];

    /*
     var currentYear = moment().year();
     var currentMonth = moment().month();

    function random(min, max) {
      return Math.floor((Math.random() * max) + min);
    }

    for (var i = 0; i < 1000; i++) {
      var start = new Date(currentYear,random(0, 11),random(1, 28),random(0, 24),random(0, 59));
      $scope.events.push({
        title: 'Event ' + i,
        type: 'warning',
        startsAt: start,
        endsAt: moment(start).add(2, 'hours').toDate()
      })
    }*/

    function showModal(action, event) {
      $modal.open({
        templateUrl: 'modalContent.html',
        controller: function($scope) {
          $scope.action = action;
          $scope.event = event;
        }
      });
    }

    // var dataObj = {
    //   id : "1012",
    //   a1 : "target",
    //   a2 : $scope.events[0].startsAt,
    //   a3 : "target3"
    // }



    function insertData1(event) {
      var res = $http.post('http://localhost:3000/reservation', {title: event.title, type: event.type, startsAt: event.startsAt, endsAt: event.endsAt});
        res.success(function(data, status, headrs, config){
          //$scope.event.title = data.a1;
        });
    }

    function deleteData(event){
      var res = $http.put('http://localhost:3000/reservation', {title: event.title, type: event.type, startsAt: event.startsAt, endsAt: event.endsAt});
        res.success(function(data, status, headrs, config){
        });
    }

    // function getData(event) {
    //   var res = $http.get('http://localhost:3000/reservations');
    //     res.success(function(data, status, headrs, config){
          


    //       event.title = data[1].id;
    //     });
    // }

    $scope.eventClicked = function(event) {
      showModal('Clicked', event);
     insertData1(event);
     //getData(event);
    };

    $scope.eventEdited = function(event) {
      //getData(event);
      showModal('Edited', event);
      insertData1(event);
    };

    $scope.eventDeleted = function(event) {
      showModal('Deleted', event);
      deleteData(event);
    };

    $scope.eventDropped = function(event) {
      showModal('Dropped', event);
    };

    $scope.toggle = function($event, field, event) {
      $event.preventDefault();
      $event.stopPropagation();
      event[field] = !event[field];
    };

  });
