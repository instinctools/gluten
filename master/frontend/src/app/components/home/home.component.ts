import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {ExecutionService} from "../../services/execution.service";
import {HomeModel} from "../../model/temp.model";

@Component({
  moduleId: module.id,
  selector: 'home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  public str: HomeModel;

  constructor(private executionService: ExecutionService, private router: Router) { }

  ngOnInit(): void {
    this.str = new HomeModel;
    this.executionService.getDefaultStr().subscribe(x => {
      this.str = x
    })
  }

}
