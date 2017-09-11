import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {AppComponent} from './components/app/app.component';
import {ExecutionComponent} from './components/execution/execution.component';
import {ExecutionService} from "./services/execution/execution.service";
import {AppRouting, appRoutingProviders} from "./app.routing";
import {HttpModule} from "@angular/http";
import {FormsModule} from "@angular/forms";
import {ResultComponent} from './components/result/result.component';
import {NodeComponent} from './components/node/node.component';
import { AlertModule } from 'ngx-bootstrap';
import {MenuComponent} from "./components/menu/menu.component";
import {ProjectComponent} from "./components/project/project.component";


@NgModule({
    imports: [
        AlertModule.forRoot(),
        BrowserModule,
        FormsModule,
        HttpModule,
        AppRouting
    ],
    exports: [

    ],
    declarations: [
        AppComponent,
        ExecutionComponent,
        ResultComponent,
        NodeComponent,
        MenuComponent,
        ProjectComponent
    ],
    providers: [
        appRoutingProviders,
        ExecutionService
    ],
    bootstrap: [AppComponent]
})
export class AppModule {
}
