import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {ResultService} from "../../services/result/result.service";
import {Result} from "../../model/result.model";
import {ExecutionService} from "../../services/execution/execution.service";

@Component({
  selector: 'app-result',
  templateUrl: './result.component.html',
  styleUrls: ['./result.component.css'],
  providers: [ResultService, ExecutionService]
})
export class ResultComponent implements OnInit {

  public results: Result[];

  constructor(private resultService: ResultService, public route: ActivatedRoute) { }

  ngOnInit() {
    this.results = [];
    this.init();
  }

  init() {
    this.route.params.subscribe(params => {
        this.resultService.getAll(params['id']).subscribe(x => {
          this.results = x;
        });
      });
  }

  convertToDate(count: number): string {
    let unixNano = new Date(count);
    return unixNano.toTimeString();
  }
}
