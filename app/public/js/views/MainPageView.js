var _ = require('underscore'),
	Backbone = require('backbone'),
	DustRenderMixins = require('../common/views/DustRenderMixins'),
	MainPageTmpl = require('../tmpl/MainPageTmpl'),
	UrlBarView = require('./UrlBarView'),
	ProgressBarView = require('../common/views/ProgressBarView');

var MainPageView = _.extend( {}, DustRenderMixins, {

	dust_template: 'MainPageTmpl.dust',

	postRender: function(html) {
		
		this.$el.html(html);

		this.$urlbar = this.$("[data-dom-ref~=url-bar]");

		new ProgressBarView({el: this.el});
		new UrlBarView({el: this.$urlbar[0]}).render();
	}

});

module.exports = Backbone.View.extend(MainPageView);