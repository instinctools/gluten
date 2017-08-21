import {NgModule} from "@angular/core";
import {RouterModule, Routes} from "@angular/router";
import {ExecutionComponent} from "./components/execution/execution.component";
import {ResultComponent} from "./components/result/result.component";

const routes: Routes = [
  {path: '', component: ExecutionComponent, pathMatch: 'full'},
  {path: 'execution', component: ExecutionComponent},
  {path: 'result/:id', component: ResultComponent}
];

export const appRoutingProviders: any[] = [];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRouting {
}
