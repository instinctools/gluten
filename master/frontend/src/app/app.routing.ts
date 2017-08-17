import {NgModule} from "@angular/core";
import {RouterModule, Routes} from "@angular/router";
import {ExecutionComponent} from "./components/execution/execution.component";
import {HomeComponent} from "./components/home/home.component";

const routes: Routes = [
  {path: '', component: HomeComponent, pathMatch: 'full'},
  {path: 'executions', component: ExecutionComponent}
];

export const appRoutingProviders: any[] = [];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRouting {
}
