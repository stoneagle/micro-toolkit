import { NgModule } from '@angular/core';
import { AppRoutingModule } from '../app-routing.module';
import { BrowserModule } from '@angular/platform-browser';
import { CommonModule  } from '@angular/common';  
import { ClarityModule  } from "@clr/angular";
import { HttpClientModule  }    from '@angular/common/http';
import { FormsModule }   from '@angular/forms';

import { SignInComponent } from './sign-in/sign-in.component';
import { ShellComponent } from './shell/shell.component';
import { MessageComponent } from './message/message.component';
import { BaseService  } from '../service/base/base.service';
import { MessageService } from '../service/base/message.service';

@NgModule({
  declarations: [
    SignInComponent,
    MessageComponent,
    ShellComponent,
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    HttpClientModule,
    FormsModule,
    ClarityModule,
  ],
  providers: [
    BaseService,
    MessageService,
  ],
  exports: [
    SignInComponent,
  ]
})
export class BaseModule { 
}
