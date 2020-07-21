import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServiceDetailsDialogComponent } from './service-details-dialog.component';

describe('ServiceDetailsDialogComponent', () => {
  let component: ServiceDetailsDialogComponent;
  let fixture: ComponentFixture<ServiceDetailsDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServiceDetailsDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServiceDetailsDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
