export class RequestInfo {
	method: string;
	headers: Record<string, string>;
	body: unknown;

	constructor(opts?: { method?: string; headers?: Record<string, string>; body?: unknown }) {
		this.method = opts?.method ?? 'POST';
		this.headers = opts?.headers ?? {};
		this.body = opts?.body ?? {};
	}

	toJSON() {
		return { method: this.method, headers: this.headers, body: this.body };
	}
}
