import { extend } from './util'
import View, { ViewAttrs } from './view'
import { write, importTemplate } from './render'
import CaptchaView from './captcha'

// Generic input form view with optional captcha support
abstract class FormView extends View<null> {
	private captcha: CaptchaView

	protected abstract send(): void

	constructor(attrs: ViewAttrs) {
		super(attrs)
		this.onClick({
			"input[name=cancel]": () =>
				this.remove(),
			".map-remove": e =>
				this.removeMapInput(e),
			".map-add": e =>
				this.addMapInput(e),
		})
		this.on("submit", e =>
			this.submit(e))

		const captcha = this.el.querySelector(".captcha-container")
		if (captcha) {
			this.captcha = new CaptchaView(captcha)
		}
	}

	// Submit form to server. Pass it to the assigned handler function
	private submit(event: Event) {
		event.preventDefault()
		this.send()
	}

	// Also destroy captcha, if any
	public remove() {
		if (this.captcha) {
			this.captcha.remove()
		}
		super.remove()
	}

	// Inject captcha data into the request struct, if any
	protected injectCaptcha(req: {}) {
		if (this.captcha) {
			extend(req, this.captcha.data())
		}
	}

	// Render a text comment about the response status below the form
	protected renderFormResponse(text: string) {
		write(() =>
			this.el.querySelector(".form-response").textContent = text)
	}

	// Load a new captcha, if present and response code is not 0
	protected reloadCaptcha() {
		if (this.captcha) {
			this.captcha.reload()
		}
	}

	// Render an additional map key-value input field pair
	protected addMapInput(event: Event) {
		write(() =>
			(event.target as Element)
				.before(this.renderKeyValuePair("", "")))
	}

	// Render a single key-value input field pair in a map subform
	private renderKeyValuePair(key: string, val: string): DocumentFragment {
		const frag = importTemplate("keyValue"),
			[k, v] = frag.querySelectorAll("input")
		k.setAttribute("value", key)
		v.setAttribute("value", val)
		return frag
	}

	// Remove a map key-vale input field pair
	protected removeMapInput(event: Event) {
		write(() =>
			(event.target as Element).closest("span").remove())
	}
}

export default FormView
