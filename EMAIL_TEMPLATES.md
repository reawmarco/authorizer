# Email Templates (HTML + Variables)

These templates are ready to use in Authorizer Dashboard (`Email Templates`).
Use Go template variables exactly as shown (`{{.variable}}`).

## Variables available

- `{{.organization.name}}`
- `{{.organization.logo}}`
- `{{.verification_url}}`
- `{{.otp}}` (OTP event only)
- `{{.user.email}}`
- `{{.user.given_name}}`
- `{{.user.family_name}}`

---

## Event: `forgot_password`
**Subject**
```text
Reset your password for {{.organization.name}}
```

**HTML**
```html
<!doctype html>
<html>
  <body style="font-family:Arial,sans-serif;background:#f6f8fb;padding:24px;">
    <table style="max-width:600px;margin:0 auto;background:#fff;border:1px solid #e5e7eb;border-radius:8px;padding:24px;">
      <tr><td style="text-align:center;"><img src="{{.organization.logo}}" alt="logo" width="36"></td></tr>
      <tr><td><h2 style="margin:16px 0 8px;">Password reset request</h2></td></tr>
      <tr><td><p>Hello {{.user.given_name}},</p></td></tr>
      <tr><td><p>We received a request to reset your password for <b>{{.user.email}}</b>.</p></td></tr>
      <tr><td style="padding:16px 0;"><a href="{{.verification_url}}" style="background:#2563eb;color:#fff;text-decoration:none;padding:10px 16px;border-radius:6px;display:inline-block;">Reset Password</a></td></tr>
      <tr><td><p>If you did not request this, you can safely ignore this email.</p></td></tr>
      <tr><td><p>Team {{.organization.name}}</p></td></tr>
    </table>
  </body>
</html>
```

## Event: `basic_auth_signup`
**Subject**
```text
Confirm your email at {{.organization.name}}
```

**HTML**
```html
<!doctype html>
<html>
  <body style="font-family:Arial,sans-serif;background:#f6f8fb;padding:24px;">
    <table style="max-width:600px;margin:0 auto;background:#fff;border:1px solid #e5e7eb;border-radius:8px;padding:24px;">
      <tr><td style="text-align:center;"><img src="{{.organization.logo}}" alt="logo" width="36"></td></tr>
      <tr><td><h2 style="margin:16px 0 8px;">Confirm your email</h2></td></tr>
      <tr><td><p>Hi {{.user.given_name}}, click below to confirm your email and activate your account.</p></td></tr>
      <tr><td style="padding:16px 0;"><a href="{{.verification_url}}" style="background:#2563eb;color:#fff;text-decoration:none;padding:10px 16px;border-radius:6px;display:inline-block;">Confirm Email</a></td></tr>
      <tr><td><p>Team {{.organization.name}}</p></td></tr>
    </table>
  </body>
</html>
```

## Event: `magic_link_login`
**Subject**
```text
Your magic login link for {{.organization.name}}
```

**HTML**
```html
<!doctype html>
<html>
  <body style="font-family:Arial,sans-serif;background:#f6f8fb;padding:24px;">
    <table style="max-width:600px;margin:0 auto;background:#fff;border:1px solid #e5e7eb;border-radius:8px;padding:24px;">
      <tr><td style="text-align:center;"><img src="{{.organization.logo}}" alt="logo" width="36"></td></tr>
      <tr><td><h2 style="margin:16px 0 8px;">Sign in with magic link</h2></td></tr>
      <tr><td><p>Hi {{.user.given_name}}, use the secure link below to sign in.</p></td></tr>
      <tr><td style="padding:16px 0;"><a href="{{.verification_url}}" style="background:#2563eb;color:#fff;text-decoration:none;padding:10px 16px;border-radius:6px;display:inline-block;">Sign In</a></td></tr>
      <tr><td><p>Team {{.organization.name}}</p></td></tr>
    </table>
  </body>
</html>
```

## Event: `update_email`
**Subject**
```text
Confirm your new email at {{.organization.name}}
```

**HTML**
```html
<!doctype html>
<html>
  <body style="font-family:Arial,sans-serif;background:#f6f8fb;padding:24px;">
    <table style="max-width:600px;margin:0 auto;background:#fff;border:1px solid #e5e7eb;border-radius:8px;padding:24px;">
      <tr><td style="text-align:center;"><img src="{{.organization.logo}}" alt="logo" width="36"></td></tr>
      <tr><td><h2 style="margin:16px 0 8px;">Confirm your new email</h2></td></tr>
      <tr><td><p>Hi {{.user.given_name}}, confirm your new email address by clicking below.</p></td></tr>
      <tr><td style="padding:16px 0;"><a href="{{.verification_url}}" style="background:#2563eb;color:#fff;text-decoration:none;padding:10px 16px;border-radius:6px;display:inline-block;">Confirm Email</a></td></tr>
      <tr><td><p>Team {{.organization.name}}</p></td></tr>
    </table>
  </body>
</html>
```

## Event: `verify_otp`
**Subject**
```text
Your OTP code for {{.organization.name}}
```

**HTML**
```html
<!doctype html>
<html>
  <body style="font-family:Arial,sans-serif;background:#f6f8fb;padding:24px;">
    <table style="max-width:600px;margin:0 auto;background:#fff;border:1px solid #e5e7eb;border-radius:8px;padding:24px;">
      <tr><td style="text-align:center;"><img src="{{.organization.logo}}" alt="logo" width="36"></td></tr>
      <tr><td><h2 style="margin:16px 0 8px;">One-time password (OTP)</h2></td></tr>
      <tr><td><p>Use the code below to continue:</p></td></tr>
      <tr><td><p style="font-size:28px;font-weight:700;letter-spacing:4px;">{{.otp}}</p></td></tr>
      <tr><td><p>This code expires soon. Do not share it.</p></td></tr>
      <tr><td><p>Team {{.organization.name}}</p></td></tr>
    </table>
  </body>
</html>
```

## Event: `invite_member`
**Subject**
```text
You are invited to join {{.organization.name}}
```

**HTML**
```html
<!doctype html>
<html>
  <body style="font-family:Arial,sans-serif;background:#f6f8fb;padding:24px;">
    <table style="max-width:600px;margin:0 auto;background:#fff;border:1px solid #e5e7eb;border-radius:8px;padding:24px;">
      <tr><td style="text-align:center;"><img src="{{.organization.logo}}" alt="logo" width="36"></td></tr>
      <tr><td><h2 style="margin:16px 0 8px;">Invitation to join</h2></td></tr>
      <tr><td><p>Hello, you were invited to join <b>{{.organization.name}}</b>.</p></td></tr>
      <tr><td style="padding:16px 0;"><a href="{{.verification_url}}" style="background:#2563eb;color:#fff;text-decoration:none;padding:10px 16px;border-radius:6px;display:inline-block;">Accept Invitation</a></td></tr>
      <tr><td><p>If you were not expecting this invitation, ignore this email.</p></td></tr>
    </table>
  </body>
</html>
```
