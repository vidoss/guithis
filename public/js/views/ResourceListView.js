var _ = require('underscore'),
	Backbone = require('backbone'),
	DustRenderMixins = require('../common/views/DustRenderMixins'),
	ResourceCollection = require('../models/ResourceListView');

var ResourceListView = _.extend({}, DustRenderMixins, {

	dust_template: 'ResourceListTmpl'

});

modules.exports = Backbone.View.extend(ResourceListView);
