import {Observable} from "rxjs/Observable";
import {Injectable} from "@angular/core";
import {Response} from "@angular/http";
import {Http} from '@angular/http';
import "rxjs/Rx";
import {Constraints} from "../../helpers/constraints";
import {HeadersService} from "../headers.service";


@Injectable()
export class ProjectService {

    private projectsURL: string;
    private buildProjectURL: string;

    constructor(private http: Http) {
        this.buildProjectURL = Constraints.baseURL + Constraints.buildProject;
        this.projectsURL = Constraints.baseURL + Constraints.projects;
    }

    getAll(): Observable<string[]> {
        return this.http.get(this.projectsURL, {headers: HeadersService.prepareHeaders()})
            .map(ProjectService.extractData)
            .catch(ProjectService.handleError)
    }

    getByKey(key: string): Observable<string> {
        return this.http
            .get(this.projectsURL + key + Constraints.separator + Constraints.editProjectByKey, {headers: HeadersService.prepareHeaders()})
            .map(ProjectService.extractData)
            .catch(ProjectService.handleError);
    }

    buildProject(json: string): Observable<any> {
        return this.http
            .post(this.buildProjectURL, json, {headers: HeadersService.prepareHeaders()})
            .map(ProjectService.extractData)
            .catch(ProjectService.handleError);
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


