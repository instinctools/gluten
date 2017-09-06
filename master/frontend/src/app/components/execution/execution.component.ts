import {Component, Input, OnInit} from '@angular/core';
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
    public dangerLabel: string;
    public successLabel: string;

    constructor(private executionService: ExecutionService, public route: ActivatedRoute,
                private router: Router) {
    }

    ngOnInit() {
        this.statusUI = "";
        this.executions = [];
        this.initData();

        this.dangerLabel = "label label-danger";
        this.successLabel = "label label-success";
    }

    showResults(id: number) {
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

    startExecution(id: number) {
        this.executionService.startExecution(id).subscribe(x => {
        });
        console.log("this ID will be start execution on slave  " + id);
    }

    stopExecution(id: number) {
        this.executionService.stopExecution(id).subscribe(x => {
        });
        console.log("this ID will be stop execution on slave  " + id);
    }

    convertToDate(id: number): string {
        let unixNano = new Date(id);
        console.log(id);
        return "DATE";
    }

    changeStatusOnClick(status: boolean): string {
        if (status) {
            return this.successLabel;
        }else {
            return this.dangerLabel;
        }
    }

}
