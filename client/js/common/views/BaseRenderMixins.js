/* Base class for all template rendering mixins */

var BaseRenderMixins = {

	/* Generates HTML string to render component */
	render: function() {
		return this;
	},

	/* Inserts HTML to DOM. Called at the end of render() */
	postRender: function(out) {
		this.$el.html(out);
	},

	/* get the data object used in render template */
	getRenderContext: function() {
		var m = this.collection || this.model ;
		return m ? m.toJSON() : {};
	},

	/* called when there is rendering error */
	onRenderError: function(err) {
		console.error(err);
	}

};

module.exports = BaseRenderMixins;