import {Observable} from "rxjs/Observable";
import {Injectable} from "@angular/core";
import {Response} from "@angular/http";
import {Http} from '@angular/http';
import "rxjs/Rx";
import {Constraints} from "../../helpers/constraints";
import {HeadersService} from "../headers.service";
import {ExecutionResult} from "../../model/execution-result.model";


@Injectable()
export class ResultService {

  private executionURL: string;

  constructor(private http: Http) {
    this.executionURL = Constraints.baseURL + Constraints.executionURL;
  }

  getAll(id: number): Observable<ExecutionResult[]> {
    return this.http
      .get(this.executionURL + id + Constraints.separator + Constraints.resultURL, {headers: HeadersService.prepareHeaders()})
      .map(ResultService.extractData)
      .catch(ResultService.handleError)
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

