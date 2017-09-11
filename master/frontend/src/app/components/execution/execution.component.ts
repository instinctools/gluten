import {Component, OnInit} from '@angular/core';
import {Execution} from "../../model/execution.model";
import {ExecutionService} from "../../services/execution/execution.service";
import {ActivatedRoute, Router} from "@angular/router";

@Component({
    moduleId: module.id,
    selector: 'execution',
    templateUrl: './execution.component.html',
    styleUrls: ['./execution.component.css']
})
export class ExecutionComponent implements OnInit {

    public executions: Execution[];
    public statusUI: string;

    //for pagination
    private offset: number;

    //for UI
    public statuses: Map<string, string>;

    constructor(private executionService: ExecutionService, public route: ActivatedRoute,
                private router: Router) {
    }

    ngOnInit() {
        this.offset = 0;

        this.statusUI = "";
        this.executions = [];
        this.initData();
        this.initStatusMap();
        this.route.params.subscribe(x => {
        })
    }

    initStatusMap() {
        this.statuses = new Map<string, string>();
        this.statuses.set("CREATED", "label label-warning");
        this.statuses.set("RUNNING", "label label-primary");
        this.statuses.set("FAILED", "label label-danger");
        this.statuses.set("COMPLETED", "label label-success");
    }

    showResults(id: string) {
        console.log(id);
        this.router.navigate(['/result', id]);
    }

    initData() {
        this.executionService.getAll(this.offset).subscribe(x => {
                this.executions = x
            }
        )
    }

    stopExecution(id: string) {
        this.executionService.stopExecution(id).subscribe(x => {
        });
        console.log("this ID will be stop execution on slave  " + id);
    }

    isValidButton(status: string): boolean {
        return status != "RUNNING";
    }

    convertToDate(unix: number): string {
        let date = new Date(unix / 1000000);
        return date.toDateString();
    }

    prevPage() {
        this.offset -= 7;
        this.initData();
    }

    nextPage() {
        this.offset += 7;
        this.initData();
    }
}
