/*
 Post action header dropdown menu
 */

const main = require('../main'),
	{$, _, Backbone, common, etc, lang, state} = main;

const MenuView = module.exports = Backbone.View.extend({
	className: 'popup-menu',
	tagName: 'ul',
	// Post menu items
	actions: ['report', 'hide'],
	events: {
		click: 'handleClick'
	},
	initialize({parent}) {
		parent._popup_menu = this;
		this.parent = parent;
		this.render();
	},
	render() {
		let html = '';
		for (let action of this.actions) {
			html += `<li data-type="${action}">${lang[action]}</li>`
		}
		this.el.innerHTML = html;
		this.parent.append(this.el);
	},
	// Forward post model to appropriate handler
	handleClick(e) {
		e.stopPropagation();
		main.request(e.target.getAttribute('data-type'), this.model);
		this.remove();
	},
	remove() {
		this.parent._popup_menu = null;
		this.el.remove();
		this.stopListening();
	}
});

main.reply('menu:extend', action =>
	_.extend(MenuView.prototype.actions, action));

main.$threads.on('click', '.control', event => {
	const {target} = event;

	// Menu already exists on the element. Remove it instead.
	if (target._popup_menu)
		return target._popup_menu.remove();

	const model = etc.getModel(target);
	if (model)
		new MenuView({parent: target, model});
});
