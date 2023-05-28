import config from "$lib/config";

export type Method = "GET" | "POST" | "PUT" | "DELETE";

export default class Requests {
  static async makeRequest(
    method: Method,
    action: string,
    body: any,
    params: any = {},
    headers: HeadersInit = {}
  ): Promise<any> {
    return this._makeRequest(method, action, params, headers, body);
  }

  static async makeAuthRequest(
    method: Method,
    action: string,
    body: any,
    params: any = {},
    headers: HeadersInit = {}
  ): Promise<any> {
    return this._makeRequest(
      method,
      action,
      params,
      {
        Authorization: localStorage.getItem("token") as string,
        ...headers,
      },
      body
    );
  }

  private static async _makeRequest(
    method: string,
    action: string,
    params: any,
    headers: HeadersInit,
    body: any
  ): Promise<any> {
    return fetch(
      `${config["backendAddress"]}/${action}?${this.parseParams(params)}`,
      {
        method: method,
        mode: "cors",
        headers: headers,
        body: body ? JSON.stringify(body) : null,
      }
    );
  }

  private static parseParams(params: any): string {
    return new URLSearchParams(params).toString();
  }
}
