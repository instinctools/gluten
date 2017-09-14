import {Observable} from "rxjs/Observable";
import {Injectable} from "@angular/core";
import {Response} from "@angular/http";
import {Http} from '@angular/http';
import "rxjs/Rx";
import {Constraints} from "../../helpers/constraints";
import {HeadersService} from "../headers.service";

@Injectable()
export class NodeService {

  private nodesURL: string;

  constructor(private http: Http) {
    this.nodesURL = Constraints.baseURL + Constraints.nodes;
  }

  getAll(): Observable<string[]> {
    return this.http.get(this.nodesURL, {headers: HeadersService.prepareHeaders()})
        .map(NodeService.extractData)
        .catch(NodeService.handleError)
  }

  private static extractData(res: Response) {
    console.log(res.json());
    return res.json();
  }

  private static handleError(error: any): Promise<any> {
    console.error('An error occurred', error);
    return Promise.reject(error.message || error);
  }
}
