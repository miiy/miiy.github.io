---
layout: base
title: Links
permalink: /links/
---

<h1 class="page-heading">{{ page.title | escape }}</h1>

{% assign links = site.data.links %}

<div class="nav-home">
  {% for group in links.groups %}
    <section class="nav-group">
      <div class="nav-group__header">
        <h2 class="nav-group__title">{{ group.title }}</h2>
        {% if group.description %}
          <p class="nav-group__description">{{ group.description }}</p>
        {% endif %}
      </div>

      <div class="nav-grid">
        {% for item in group.items %}
          {% assign host = item.url | remove: 'https://' | remove: 'http://' | split: '/' | first %}
          <a class="nav-card" href="{{ item.url }}" title="{{ item.desc | default: item.name }}" target="_blank" rel="noopener noreferrer">
            <span class="nav-card__icon" aria-hidden="true">
              <span>{{ item.name | remove: "【" | remove: "】" | strip | slice: 0, 1 | upcase }}</span>
              {% if item.icon %}
                <img src="{{ item.icon }}" alt="" loading="lazy" onerror="this.remove()">
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
