var _ = require('underscore'),
	Backbone = require('backbone'),
	DustRenderMixins = require('../common/views/DustRenderMixins'),
	ResourceCollection = require('../models/ResourceListView');

var ResourceFilterView = _.extend({}, DustRenderMixins, {

	dust_template: 'ResourceFilterTmpl'

});

modules.exports = Backbone.View.extend(ResourceFilterView);