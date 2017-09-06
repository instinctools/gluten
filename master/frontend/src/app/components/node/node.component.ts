import {Component, OnInit} from '@angular/core';
import {NodeService} from "../../services/node/node.service";
import {ActivatedRoute} from "@angular/router";

@Component({
    selector: 'app-node',
    templateUrl: './node.component.html',
    styleUrls: ['./node.component.css'],
    providers: [NodeService]
})
export class NodeComponent implements OnInit {

    public nodes: string[];
    public iterator: number;

    constructor(private nodeService: NodeService, public route: ActivatedRoute) {
    }

    ngOnInit() {
        this.iterator = 1;
        this.nodes = [];
        this.init();

    }

    init() {
        this.nodeService.getAll().subscribe(x => {
            this.nodes = x
        });
    }



}
