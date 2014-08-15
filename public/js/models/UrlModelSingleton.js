var _ = require('underscore'),
	Backbone = require('backbone'),
	AppDispatcher = require('../dispatcher/AppDispatcher'),
	AppConstants = require('../constants/AppConstants');

var UrlModel = Backbone.Model.extend({

	initialize: function() {
		this.listenTo(AppDispatcher, AppConstants.URL_FETCH, this.fetch);
	}
	
});

var _instance = null;

module.exports = {
	getInstance: function() {
		return _instance ? _instance : (_instance = new UrlModel())
	}
};