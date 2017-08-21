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

  constructor(private executionService: ExecutionService, public route: ActivatedRoute,
              private router: Router) {
  }

  ngOnInit() {
    this.executions = [];
    this.initData();
  }

  showResults(id: number) {
    this.router.navigate(['/result', id]);
  }

  initData() {
    this.executionService.getAll().subscribe(x => {
        this.executions = x
      }
    )
  }

  startExecution(id: number) {
    this.executionService.startExecution(id).subscribe(x => {});
    console.log("this ID will be start execution on slave  " + id);
  }

  stopExecution(id: number) {
    this.executionService.stopExecution(id).subscribe(x => {});
    console.log("this ID will be stop execution on slave  " + id);
  }

}
