import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {ProjectService} from "../../services/project/project.service";

@Component({
    moduleId: module.id,
    selector: 'project',
    templateUrl: './project.component.html',
    styleUrls: ['./project.component.css'],
    providers: [ProjectService]

})
export class ProjectComponent implements OnInit {

    public project: string;
    public projects: string[];

    constructor(private projectService: ProjectService, public route: ActivatedRoute,
                private router: Router) {
    }

    ngOnInit() {
        this.project = "";
        this.projects = [];
        this.initListProjects();
    }

    buildProject(json: string) {
        this.projectService.runProject(json).subscribe(x => {

        });

        this.router.navigate(['/execution']);
    }

    initListProjects() {
        this.projectService.getAll().subscribe(x => {
            this.projects = x;
        })
    }

    editStructure(key: string) {
        this.projectService.getByKey(key).subscribe(x => {
            this.project = x;
        })
    }

}
