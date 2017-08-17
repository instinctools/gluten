import {Component, Input, OnInit} from '@angular/core';
import {Execution} from "../../model/execution.model";
import {ExecutionService} from "../../services/execution.service";
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

  initData() {
    this.executionService.getAll().subscribe(x => {
        this.executions = x
      }
    )
  }

  stopExecution(id: number) {
    console.log("this ID will be stop execution on slave" + id);
  }

}
