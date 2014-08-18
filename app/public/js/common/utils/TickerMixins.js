var _ = require('underscore');

var TickerMixins = {

	tick_interval: 2000, /* milliseconds */

	tickStart: function() {
		this.timer = setInterval(_.bind(this.tick, this), this.tick_interval);
	},

	tickStop: function() {
		clearInterval(this.timer);
	},

	tick: function() {
	}
};

module.exports = TickerMixins;