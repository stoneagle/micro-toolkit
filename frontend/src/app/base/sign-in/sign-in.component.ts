import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import { AppConfig } from '../../service/app.config';
import { BaseService  } from '../../service/base/base.service';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent implements OnInit {
  CompanyName: string;
  username: string;
  password: string;
  // atype: string = "admin";

  constructor(
    private baseService: BaseService,
    private router: Router
  ) { }

  ngOnInit() {
    this.CompanyName = AppConfig.settings.app.company;
  }

  submit() {
    this.baseService.login(this.username, this.password)
    .subscribe(res => {
      this.router.navigate(['/toolkit']);
    })
  }
}
