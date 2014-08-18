var _ = require('underscore'),
	Nanobar = require('nanobar'),
	Backbone = require('backbone'),
	TickerMixins = require('../utils/TickerMixins');

var ProgressBarView = _.extend( {}, TickerMixins, {

	tick_interval: 100,
	tick_increment: 10,
	tick_max: 100,

	options: {
		bg: "#acf"
	},

	initialize: function(options) {
		
		_.extend(this.options, options);
		this.nanobar = new Nanobar(options);
		
		Backbone.$.ajaxSetup({
			beforeSend: _.bind(this.ajaxStart, this),
			complete: _.bind(this.ajaxEnd, this)
		});

		this.progress = 0;
	},

	ajaxStart: function() {
		this.nanobar.go(0);
		this.tickStart();
	},

	tick: function() {
		this.progress += this.tick_increment;
		(this.progress == this.tick_max) ? this.tickStop() : this.nanobar.go(this.progress);
	},

	ajaxEnd: function() {
		this.tickStop();
		this.progress = 0;
		this.nanobar.go(100);
	}
});

module.exports = Backbone.View.extend( ProgressBarView );