import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppComponent } from './components/app/app.component';
import { ExecutionComponent } from './components/execution/execution.component';
import {ExecutionService} from "./services/execution.service";
import {AppRouting, appRoutingProviders} from "./app.routing";
import {HttpModule} from "@angular/http";
import {FormsModule} from "@angular/forms";
import { HomeComponent } from './components/home/home.component';


@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRouting
  ],
  declarations: [
    AppComponent,
    ExecutionComponent,
    HomeComponent
  ],
  providers: [
    appRoutingProviders,
    ExecutionService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
