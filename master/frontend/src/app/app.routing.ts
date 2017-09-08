import {NgModule} from "@angular/core";
import {RouterModule, Routes} from "@angular/router";
import {ExecutionComponent} from "./components/execution/execution.component";
import {ResultComponent} from "./components/result/result.component";
import {NodeComponent} from "./components/node/node.component";
import {MenuComponent} from "./components/menu/menu.component";
import {ProjectComponent} from "./components/project/project.component";

const routes: Routes = [
    {path: '', component: ExecutionComponent, pathMatch: 'full'},
    {path: 'menu', component: MenuComponent},
    {path: 'project', component: ProjectComponent},
    {path: 'execution', component: ExecutionComponent},
    {path: 'node', component: NodeComponent},
    {path: 'result/:id', component: ResultComponent}
];

export const appRoutingProviders: any[] = [];


@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRouting {
}
