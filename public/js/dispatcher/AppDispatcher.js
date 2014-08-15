var _ = require('underscore'),
	Dispatcher = require('../common/dispatcher/Dispatcher');

var AppDispatcher = _.extend( {}, Dispatcher);

module.exports = AppDispatcher;