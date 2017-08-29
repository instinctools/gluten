import {Headers} from "@angular/http";

export class HeadersService {

  public static prepareHeaders() {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    return headers;
  }
}
