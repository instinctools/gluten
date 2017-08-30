import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {ResultService} from "../../services/result/result.service";
import {ExecutionResult} from "../../model/execution-result.model";
import {ExecutionService} from "../../services/execution/execution.service";

@Component({
  selector: 'app-result',
  templateUrl: './result.component.html',
  styleUrls: ['./result.component.css'],
  providers: [ResultService, ExecutionService]
})
export class ResultComponent implements OnInit {

  public results: ExecutionResult[];

  constructor(private resultService: ResultService, public route: ActivatedRoute) { }

  ngOnInit() {
    this.results = [];
    this.init();
  }

  init() {
    this.results = [];

    this.route.params.subscribe(params => {
        this.resultService.getAll(+params['id']).subscribe(x => {
          this.results = x;
        });
      });
  }

}
