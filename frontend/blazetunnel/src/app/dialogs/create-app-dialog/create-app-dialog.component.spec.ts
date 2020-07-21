import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateAppDialogComponent } from './create-app-dialog.component';

describe('CreateAppDialogComponent', () => {
  let component: CreateAppDialogComponent;
  let fixture: ComponentFixture<CreateAppDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CreateAppDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateAppDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
