import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes  } from '@angular/router';
import { AutobuildComponent } from './controller/storybox/autobuild/autobuild.component';
import { SignInComponent } from './base/sign-in/sign-in.component';
import { ShellComponent } from './base/shell/shell.component';

const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'login', component: SignInComponent  },
  { path: 'toolkit', component: ShellComponent, 
      children:[
        { path: 'autobuild', component: AutobuildComponent  },
      ]
  },
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }
