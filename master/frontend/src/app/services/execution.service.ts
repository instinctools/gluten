import { Observable } from "rxjs/Observable";
import {Injectable} from "@angular/core";
import { Response } from "@angular/http";
import { Http } from '@angular/http';
import "rxjs/Rx";
import {Constraints} from "../helpers/constraints";
import {Execution} from "../model/execution.model";
import {HeadersService} from "./headers.service";
import {HomeModel} from "../model/temp.model";


@Injectable()
export class ExecutionService {

  private executionURL: string;

  constructor(private http: Http) {
    this.executionURL = Constraints.baseURL + Constraints.executionURL;
  }

  getAll(): Observable<Execution[]> {
    return this.http.get(this.executionURL, { headers: HeadersService.prepareHeaders() })
      .map(this.extractData)
      .catch(this.handleError)
  }

  getDefaultStr(): Observable<HomeModel> {
    return this.http.get(Constraints.baseURL, { headers: HeadersService.prepareHeaders() })
      .map(this.extractData)
      .catch(this.handleError)
  }


  private extractData(res: Response) {
    console.log(res.json());
    return res.json();
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error);
    return Promise.reject(error.message || error);
  }

}
