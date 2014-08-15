
var $  = require('jquery'),
	_  = require('underscore'),
	BaseRenderMixins = require('./BaseRenderMixins');

// Base backbone view to render dust templates

var DustRenderMixins = _.extend( {}, BaseRenderMixins, {

	dust_template: '',

	render: function() {
		
		if (!this.dust_template) {
			console.error("dust_template not specified in view");
			return this;
		}

		dust.render(this.dust_template, this.getRenderContext(), _.bind(function(err, out) {
			if (err) {
				this.onRenderError(err);
				return;
			}

			this.postRender(out);

		}, this));

		return this;
	}

});

module.exports = DustRenderMixins;