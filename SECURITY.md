# Security Policy

## Supported Versions

I actively maintain and provide security updates for the following versions of Passgen:

| Version | Supported |
|---------|-----------|
| v0.4.0  | ✅ Yes |
| v0.3.0  | ✅ Yes |
| v0.2.0  | ✅ Yes |
| v0.1.0  | ✅ Yes |

If you are using an unsupported version, please upgrade to the latest stable release before reporting a vulnerability.

---

## Reporting a Vulnerability

If you believe you have found a security vulnerability in Passgen, please report it responsibly and privately.

📧 Security Contact Email  
bilalbaraz@windowslive.com

Please **do not create public GitHub issues** for suspected vulnerabilities.

---

## What Is Considered a Vulnerability

The following are examples of security vulnerabilities in Passgen:

- Cryptographic implementation flaws
- Exposure of generated passwords in logs, memory, or clipboard unexpectedly
- Timing attacks or side-channel leaks
- Dependency vulnerabilities affecting cryptographic or entropy sources
- Privilege escalation or sandbox escape risks
- Remote code execution or command injection risks

The following are typically **not considered vulnerabilities**:

- Feature requests
- Usage questions
- Non-security related bugs
- Issues in unsupported versions

---

## Disclosure Process

We follow a **coordinated vulnerability disclosure** process:

1. Reporter submits vulnerability privately
2. I acknowledge receipt within **7 days**
3. I investigate and validate the vulnerability
4. I develop and test a fix
5. I release a patched version
6. Public disclosure may happen after fix release

---

## Disclosure Timeline Expectations

- Initial response: within **7 days**
- Investigation and fix target: within **30–90 days** depending on severity
- Critical vulnerabilities may be fixed faster

These timelines may vary depending on complexity and impact.

---

## Safe Harbor

I support responsible security research.  
I will not take legal action against researchers who:

- Act in good faith
- Avoid privacy violations
- Avoid service disruption
- Give me reasonable time to fix the vulnerability before disclosure

---

## Security Best Practices for Users

I recommend users:

- Always use the latest version
- Avoid storing generated passwords in plaintext files
- Use secure clipboard managers if clipboard integration is used
- Verify binary releases via checksums or signatures if provided

---

## Encryption / Secure Communication (Optional)

If you prefer encrypted communication, you may request my PGP public key via the security email.

---

## Thank You

I appreciate responsible disclosure and security research efforts that help improve Passgen security.
