var AppDispatcher = require('../dispatcher/AppDispatcher'),
	AppConstants = require('../constants/AppConstants');

var AppActions = {

	urlEntered: function(url) {
		AppDispatcher.trigger(AppConstants.URL_ENTERED, url);
	},

	fetchUrl: function(options) {
		AppDispatcher.trigger(AppConstants.URL_FETCH, options);
	}

};

module.exports = AppActions;