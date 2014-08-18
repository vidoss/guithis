var _ = require('underscore'),
	Backbone = require('backbone'),
	DustRenderMixins = require('../common/views/DustRenderMixins'),
	UrlBarTmpl = require('../tmpl/UrlBarTmpl'),
	AppActions = require('../actions/AppActions'),
	UrlModelSingleton = require('../models/UrlModelSingleton');

var UrlBarView = _.extend( {}, DustRenderMixins, {

	dust_template: "UrlBarTmpl.dust",

	events: {
		"submit [data-dom-ref~=url-form]"  : "onUrlSubmit",
		"keydown [data-dom-ref~=url-input]": "clearError"
	},

	initialize: function() {

		this.model = UrlModelSingleton.getInstance();
		
		this.listenTo(this.model, 'error', this.onFetchError);
	},

	postRender: function(html) {
		this.$el.html(html);
		
		this.$url = this.$("[data-dom-ref~=url-input]");
		this.$error = this.$("[data-dom-ref~=url-error]");
	},

	onUrlSubmit: function() {

		try {
			AppActions.fetchUrl({url: this.$url.val()});
		} catch(e) {
			console.error(e);
		}
		return false;
	},

	onFetchError: function(model, resp, options) {
		console.error(resp);
		this.$error.find('h1').text(resp.status);
		this.$error.find('pre').text(resp.responseText);
		this.$error.removeClass("hidden");
	},

	clearError: function() {
		this.$error.addClass("hidden");
	}

});

module.exports = Backbone.View.extend(UrlBarView);
