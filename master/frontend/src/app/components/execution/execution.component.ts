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

    //for UI
    public statuses: Map<string, string>;

    constructor(private executionService: ExecutionService, public route: ActivatedRoute,
                private router: Router) {
    }

    ngOnInit() {
        this.statusUI = "";
        this.executions = [];
        this.initData();
        this.initStatusMap();
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

    showNodes() {
        this.router.navigate(['/node']);
    }

    initData() {
        this.executionService.getAll().subscribe(x => {
                this.executions = x
            }
        )
    }

    startExecution(id: string) {
        this.executionService.startExecution(id).subscribe(x => {
        });
        console.log("this ID will be start execution on slave  " + id);
    }

    stopExecution(id: string) {
        this.executionService.stopExecution(id).subscribe(x => {
        });
        console.log("this ID will be stop execution on slave  " + id);
    }

    convertToDate(count: number): string {
        let unixNano = new Date(count);
        return unixNano.toTimeString();
    }
}
