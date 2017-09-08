import {Observable} from "rxjs/Observable";
import {Injectable} from "@angular/core";
import {Response} from "@angular/http";
import {Http} from '@angular/http';
import "rxjs/Rx";
import {Constraints} from "../../helpers/constraints";
import {Execution} from "../../model/execution.model";
import {HeadersService} from "../headers.service";


@Injectable()
export class ExecutionService {

    private executionURL: string;
    public buildProjectURL: string;

    constructor(private http: Http) {
        this.executionURL = Constraints.baseURL + Constraints.executions;
        this.buildProjectURL= Constraints.baseURL + Constraints.buildProject;
    }

    getAll(offset: number): Observable<Execution[]> {
        return this.http
            .post(this.executionURL, offset, {headers: HeadersService.prepareHeaders()})
            .map(ExecutionService.extractData)
            .catch(ExecutionService.handleError)
    }

    stopExecution(id: string): Observable<any> {
        return this.http
            .post(this.executionURL + id + Constraints.separator + Constraints.stopExecution, {headers: HeadersService.prepareHeaders()})
            .map(ExecutionService.extractData)
            .catch(ExecutionService.handleError);
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
