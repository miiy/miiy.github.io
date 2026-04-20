---
# Feel free to add content and custom Front Matter to this file.
# To modify the layout, see https://jekyllrb.com/docs/themes/#overriding-theme-defaults

layout: base
title: Home
---

{% assign nav = site.data.navigation %}

<div class="nav-home">
  <section class="nav-home__hero">
    <h1 class="nav-home__title">Quick Access</h1>
  </section>

  {% for group in nav.groups %}
    <section class="nav-group">
      <div class="nav-group__header">
        <div>
          <h2 class="nav-group__title">{{ group.title }}</h2>
          {% if group.description %}
            <p class="nav-group__description">{{ group.description }}</p>
          {% endif %}
        </div>
        <span class="nav-group__count">{{ group.items | size }} 个站点</span>
      </div>

      <div class="nav-grid">
        {% for item in group.items %}
          {% assign host = item.url | remove: 'https://' | remove: 'http://' | split: '/' | first %}
          <a class="nav-card" href="{{ item.url }}" target="_blank" rel="noopener noreferrer">
            <span class="nav-card__icon" aria-hidden="true">
              {% if item.icon %}
                <img src="{{ item.icon }}" alt="" loading="lazy">
              {% else %}
                <span>{{ item.name | slice: 0, 1 | upcase }}</span>
              {% endif %}
            </span>

            <span class="nav-card__content">
              <span class="nav-card__name">{{ item.name }}</span>
              <span class="nav-card__host">{{ host }}</span>
              {% if item.desc %}
                <span class="nav-card__desc">{{ item.desc }}</span>
              {% endif %}
            </span>
          </a>
        {% endfor %}
      </div>
    </section>
  {% endfor %}
</div>
