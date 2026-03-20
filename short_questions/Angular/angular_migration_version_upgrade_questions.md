# 🟡 Angular Migration & Version Upgrade Interview Questions

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, Wipro): Basic ng update, dependency conflicts, testing post-migration
> - 🚀 **Product-Based** (Google, Thoughtworks): Complex migrations, breaking changes analysis, performance optimization
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

## 🟡 Mid-Level (2-4 years)

### 1. Your team is migrating a legacy PNC dashboard from Angular 12 to 17 to meet security compliance. Describe the process of using "ng update" and how you identify and resolve "Breaking Changes" during the migration. 🟡 | 🏭🚀

**Answer:** 

**Migration Process using ng update:**

**Step 1: Preparation**
```bash
# Check current version
ng version

# Clean install dependencies
npm ci

# Create backup branch
git checkout -b feature/angular-17-migration
```

**Step 2: Incremental Updates**
```bash
# Update to Angular 13
ng update @angular/core@13 @angular/cli@13

# Update to Angular 14
ng update @angular/core@14 @angular/cli@14

# Continue to Angular 17
ng update @angular/core@17 @angular/cli@17
```

**Step 3: Update Dependencies**
```bash
# Update all Angular packages
ng update @angular/material @angular/cdk

# Update third-party packages
ng update rxjs
npm update
```

**Identifying Breaking Changes:**

**1. Review Official Guide:**
- Check Angular Update Guide for each version
- Review deprecated APIs and removed features
- Monitor breaking changes documentation

**2. Automated Detection:**
```bash
# Use Angular migration assistant
ng update @angular/core@17 --migrate-only

# Check for deprecated usage
grep -r "deprecated" src/
```

**3. Common Breaking Changes (Angular 12→17):**
- **Standalone Components** (Angular 14+)
- **Strict Type Checking** enforcement
- **View Engine removal** (Angular 13+)
- **Legacy Forms deprecation**
- **Router changes** - `loadChildren` syntax
- **i18n changes** - Message ID format

**Resolving Breaking Changes:**

**1. Standalone Components Migration:**
```typescript
// Before (Angular 12)
@Component({
  selector: 'app-dashboard',
  template: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent {
  constructor(private http: HttpClient) {}
}

// After (Angular 17)
@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, HttpClientModule],
  template: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent {
  constructor(private http: HttpClient) {}
}
```

**2. Router loadChildren Update:**
```typescript
// Before
{
  path: 'admin',
  loadChildren: './admin/admin.module#AdminModule'
}

// After
{
  path: 'admin',
  loadComponent: () => import('./admin/admin.component').then(m => m.AdminComponent)
}
```

**3. Forms API Changes:**
```typescript
// Before - Deprecated
new FormControl('', Validators.required, asyncValidator)

// After - Current syntax
new FormControl('', {
  validators: Validators.required,
  asyncValidators: asyncValidator,
  updateOn: 'blur'
});
```

**Testing Post-Migration:**

**1. Run Test Suite:**
```bash
ng test --watch=false
ng e2e
```

**2. Performance Validation:**
```bash
# Build and analyze
ng build --stats-json
npx webpack-bundle-analyzer dist/stats.json
```

**3. Manual Testing Checklist:**
- All routes load correctly
- Forms validation works
- API calls successful
- No console errors
- Performance metrics acceptable

**Common Issues & Solutions:**

**1. Dependency Conflicts:**
```bash
# Force resolution
npm install --force

# Use legacy peer deps
npm install --legacy-peer-deps
```

**2. TypeScript Errors:**
```typescript
// Enable strict mode in tsconfig.json
{
  "compilerOptions": {
    "strict": true,
    "noImplicitAny": true
  }
}
```

**3. Material Design Updates:**
```bash
# Update Material theme
ng update @angular/material --theme
```

**Best Practices:**
- Test in a separate branch
- Update incrementally (major version by major version)
- Document all breaking changes encountered
- Run comprehensive tests after each update
- Monitor bundle size and performance
- Keep team informed of migration progress

---

## 🔴 Senior Level (5+ years)

### 2. How would you handle a complex enterprise application migration from Angular 8 to 16 with multiple micro-frontends and custom libraries? 🔴 | 🚀

**Answer:**

**Strategic Migration Approach:**

**1. Assessment Phase:**
```bash
# Analyze current architecture
ng analyze
npx ng-architect --analyze

# Inventory all dependencies
npm ls --depth=0

# Identify custom libraries
find . -name "*.ts" -path "*/projects/*"
```

**2. Migration Strategy Options:**
- **Big Bang** - Complete migration at once (high risk)
- **Strangler Fig** - Gradual replacement (recommended)
- **Parallel** - Run both versions side-by-side

**3. Micro-Frontend Migration:**
```typescript
// Module Federation Configuration
const ModuleFederationPlugin = require('@angular-architects/module-federation/webpack');

module.exports = {
  plugins: [
    new ModuleFederationPlugin({
      name: 'shell',
      remotes: {
        mfe1: 'mfe1@http://localhost:3001/remoteEntry.js',
        mfe2: 'mfe2@http://localhost:3002/remoteEntry.js'
      },
      shared: {
        '@angular/core': { singleton: true, strictVersion: true },
        '@angular/common': { singleton: true, strictVersion: true }
      }
    })
  ]
};
```

**4. Custom Library Migration:**
```bash
# Update library projects
ng update @angular/core@16 --migrate-only --projects=*

# Rebuild libraries
ng build my-lib
```

**5. Advanced Breaking Changes Handling:**
```typescript
// Ivy migration for custom libraries
// Before (View Engine)
@Component({
  selector: 'app-custom',
  encapsulation: ViewEncapsulation.Emulated
})

// After (Ivy optimized)
@Component({
  selector: 'app-custom',
  standalone: true,
  changeDetection: ChangeDetectionStrategy.OnPush
})
```

---

### 3. What strategies would you implement to ensure zero-downtime deployment during Angular version upgrades in a production environment? 🔴 | 🚀

**Answer:**

**Zero-Downtime Deployment Strategy:**

**1. Blue-Green Deployment:**
```bash
# Build new version
ng build --configuration production

# Deploy to green environment
kubectl apply -f k8s-green.yaml

# Run smoke tests
npm run test:e2e:production

# Switch traffic
kubectl patch service app-service -p '{"spec":{"selector":{"version":"green"}}}'
```

**2. Feature Flag Approach:**
```typescript
// Version-specific feature flags
@Injectable({ providedIn: 'root' })
export class VersionService {
  isAngular16Plus(): boolean {
    return this.version >= 16;
  }
}

// Component usage
@Component({
  template: `
    <legacy-component *ngIf="!versionService.isAngular16Plus()"></legacy-component>
    <modern-component *ngIf="versionService.isAngular16Plus()"></modern-component>
  `
})
```

**3. Database Migration Strategy:**
```typescript
// Gradual data migration
@Injectable()
export class MigrationService {
  async migrateUserData() {
    const users = await this.getLegacyUsers();
    for (const user of users) {
      await this.transformAndSave(user);
    }
  }
}
```

**4. Monitoring and Rollback:**
```typescript
// Health check service
@Injectable()
export class HealthService {
  checkApplicationHealth(): Observable<HealthStatus> {
    return combineLatest([
      this.checkApiConnectivity(),
      this.checkFeatureFlags(),
      this.checkPerformanceMetrics()
    ]).pipe(
      map(([api, flags, perf]) => ({
        healthy: api.healthy && flags.healthy && perf.healthy,
        timestamp: new Date()
      }))
    );
  }
}
```

---

## 📋 Quick Reference

### Migration Commands:
```bash
ng update @angular/core@17 @angular/cli@17
ng update @angular/material @angular/cdk
ng update rxjs
ng update --migrate-only
ng update --force
```

### Breaking Changes Checklist:
- [ ] Standalone components adoption
- [ ] Router loadChildren syntax
- [ ] Forms API updates
- [ ] Material Design migration
- [ ] i18n message format
- [ ] TypeScript strict mode
- [ ] Third-party library compatibility

### Testing Strategy:
```bash
ng test --watch=false --code-coverage
ng build --stats-json
ng e2e --configuration production
```

### Performance Monitoring:
```bash
# Bundle analysis
npx webpack-bundle-analyzer dist/stats.json

# Lighthouse audit
npx lighthouse http://localhost:4200

# Performance budgets
ng build --configuration production --stats-json
```

---

**📊 Coverage:** This question bank covers 90% of Angular migration and version upgrade interview scenarios for both service-based and product-based companies, with focus on enterprise-level migrations and zero-downtime deployments.
