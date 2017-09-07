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
//angular-material modules
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {NoopAnimationsModule} from '@angular/platform-browser/animations';
import {MdButtonModule, MdCheckboxModule} from '@angular/material';


@NgModule({
    imports: [
        MdButtonModule, MdCheckboxModule,
        NoopAnimationsModule,
        BrowserAnimationsModule,
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
        NodeComponent
    ],
    providers: [
        appRoutingProviders,
        ExecutionService
    ],
    bootstrap: [AppComponent]
})
export class AppModule {
}
