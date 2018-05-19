import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AutobuildComponent } from './autobuild.component';

describe('AutobuildComponent', () => {
  let component: AutobuildComponent;
  let fixture: ComponentFixture<AutobuildComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AutobuildComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AutobuildComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
