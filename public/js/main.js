var $ = require('jquery'),
	Backbone = require('backbone'),
	MainPageView = require('./views/MainPageView');

Backbone.$  = $; // set bb $ to jQuery

$(function(){
	new MainPageView({el: document.body}).render();
});
