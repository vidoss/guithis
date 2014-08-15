var _ = require('underscore'),
	Backbone = require('backbone'),
	DustRenderMixins = require('../common/views/DustRenderMixins'),
	ResourceCollection = require('../models/ResourceListView');

var ResourceItemView = _.extend({}, DustRenderMixins, {

	dust_template: 'ResourceItemTmpl'

});

modules.exports = Backbone.View.extend( ResourceItemView );