# 🏢 Angular Interview Questions — Service-Based Companies

> Companies like **TCS, Infosys, Wipro, Cognizant, Capgemini, Accenture, HCL, Tech Mahindra**, etc.

Service-based companies focus on **practical Angular skills** — building SPAs, working with templates and forms, integrating REST APIs, and understanding core Angular concepts. Interviews are **moderate in depth** and emphasize real-world usage over deep internals.

---

## 📂 Category Index

| # | File | Topic | Difficulty |
|---|------|--------|------------|
| 01 | [Angular Basics & Components](./01_angular_basics_and_components.md) | Components, lifecycle hooks, data binding, decorators | 🟢 Easy |
| 02 | [Directives, Services & DI](./02_directives_services_di.md) | Built-in directives, custom directives, services, dependency injection | 🟢 Easy–Medium |
| 03 | [Routing & Navigation](./03_routing_navigation.md) | RouterModule, route guards, lazy loading, ActivatedRoute | 🟡 Medium |
| 04 | [Forms](./04_forms.md) | Template-driven forms, reactive forms, validators, FormBuilder | 🟡 Medium |
| 05 | [HTTP & RxJS](./05_http_rxjs.md) | HttpClient, Observables, common RxJS operators, error handling | 🟡 Medium |
| 06 | [Pipes, Modules & Change Detection](./06_pipes_modules_change_detection.md) | Built-in pipes, custom pipes, NgModule, change detection basics | 🟡 Medium |
| 07 | [State Management & NgRx](./07_state_management_ngrx.md) | Component state, service-based state, intro to NgRx | 🟡 Medium |
| 08 | [Testing & DevOps](./08_testing_and_devops.md) | Unit testing with Jasmine/Karma, Angular CLI build, CI/CD basics | 🟡 Medium |

---

## 🎯 Interview Focus Areas (Service Companies)

- ✅ Angular fundamentals: components, modules, lifecycle hooks, decorators
- ✅ Data binding: interpolation, property binding, event binding, two-way binding
- ✅ Directives: `*ngIf`, `*ngFor`, `*ngSwitch`, custom attribute directives
- ✅ Services & DI: `@Injectable`, `providedIn: 'root'`, constructor injection
- ✅ Routing: `RouterModule`, `routerLink`, route parameters, route guards
- ✅ Forms: template-driven vs reactive, validators, `FormGroup`, `FormControl`
- ✅ HTTP: `HttpClient`, `HttpInterceptor`, error handling with `catchError`
- ✅ RxJS basics: `Observable`, `subscribe`, `map`, `filter`, `switchMap`
- ✅ Pipes: `DatePipe`, `CurrencyPipe`, custom pipes, `async` pipe

---

## 💡 Tips for Service Company Angular Interviews

1. **Know the component lifecycle** — `ngOnInit`, `ngOnDestroy`, `ngOnChanges` are always asked
2. **`*ngIf` vs `[hidden]`** — understand DOM removal vs CSS visibility trade-offs
3. **Template-driven vs Reactive forms** — know when to use each and the key differences
4. **HttpClient basics** — HTTP GET/POST, setting headers, interceptors for auth tokens
5. **Observable vs Promise** — be clear on the conceptual difference and when to use each
6. **Lazy loading modules** — always mention this as a performance optimization in routing questions
7. **Custom pipes** — know how to implement `PipeTransform` with a practical example
8. **`@Input()` / `@Output()`** — component communication is frequently tested
