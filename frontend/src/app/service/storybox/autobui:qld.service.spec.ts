import { TestBed, inject } from '@angular/core/testing';

import { AutobuildService } from './autobuild.service';

describe('AutobuildService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AutobuildService]
    });
  });

  it('should be created', inject([AutobuildService], (service: AutobuildService) => {
    expect(service).toBeTruthy();
  }));
});
