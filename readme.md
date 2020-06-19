# Html

A highly opinionated code driven HTML generator.

Usage:
```golang
    page := []HtmlElement{
		DocType(),
		Html(Lang("en")).AddElements(
			Head().AddElements(
				Meta(Key("charset", "utf-8")),
				Title().AddElements(Content("Page Title")),
				Style().AddElements(Content(
					"ul.samples {list-style: none; margin: 0; padding: 0;" +
					" margin-block-start: 1px; margin-block-end: 1px;}\n" +
					"ul.samples li {display: inline-block; width: 300px;}\n" +
					"ul.samples li.good {background-color: #008000;}\n")),
			),
			Body().AddElements(
				Heading(Level(1)).AddElements(Content("The Title of the Page")),
				P().AddElements(Content("This is the first sentence."))),
			Ul().AddElements(
				Li().AddElements(
					Img(Src("data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAAAAABXZoBIAAAA7UlEQVR4"+
						"nGJiwAMGXJJZUFCwrmO95LL/3&#43;tBfBaYhByblY1AMIj1ZHLg54sHQSxGqJzBPn4o61/SV4Zn728imyh4&#43;"+
						"y8IHNv2/SMWywPmZP/9e5abQXsWNqfxMcz6G4XLtZ8YPjKkMOL0Fve&#43;v244JRmUPz5ckINTNvDD37/lErhkdX"+
						"f9/TtNGpesQOyfv7txW/zz708HKJMFVUYvxJSF4dpBbJrUpzz9&#43;/fvr21YpCSK7oKC96QfppS401VwyAdiBpL"+
						"ganCsHA7gwJAyX/MIJPWllRtVHOzawEAGhuub//Z8wO1BTEB&#43;0gQEAAD//3QpXYUSeN4MAAAAAElFTkSuQmCC")),
				),
			),
		),
	}

	output := ""
	for _, el := range page {
		output += el.Render()
	}
```

HTML5 elements are currently supported although they are mostly
only at the bare minumum.

Notes:
This is not optimized yet, first we'll work on correctness, then once it's fairly stable then working on making it 
faster will happen.

All valid html5 tags *want* to be implemented, but for now it's driven by the ones currently in use. 

TODO:
 * Add Validation
 * Add Variables & Conditionals
 * Optimize
 * Test
 * Add CSS

HTML Tags
```
<!DOCTYPE>
<a>
<abbr>
<address>
<area>
<article>
<aside>
<audio>
<b>
<base>
<bdi>
<bdo>
<blockquote>
<body>
<br>
<button>
<canvas>
<caption>
<cite>
<code>
<col>
<colgroup>
<data>
<datalist>
<dd>
<del>
<details>
<dfn>
<dialog>
<div>
<dl>
<dt>
<em>
<embed>
<fieldset>
<figcaption>
<figure>
<footer>
<form>
<h1> - <h6>
<head>
<header>
<hr>
<html>
<i>
<iframe>
<img>
<input>
<ins>
<kbd>
<label>
<legend>
<li>
<link>
<main>
<map>
<mark>
<meta>
<meter>
<nav>
<noscript>
<object>
<ol>
<optgroup>
<option>
<output>
<p>
<param>
<picture>
<pre>
<progress>
<q>
<rp>
<rt>
<ruby>
<s>
<samp>
<script>
<section>
<select>
<small>
<source>
<span>
<strong>
<style>
<sub>
<summary>
<sup>
<svg>
<table>
<tbody>
<td>
<template>
<textarea>
<tfoot>
<th>
<thead>
<time>
<title>
<tr>
<track>
<u>
<ul>
<var>
<video>
<wbr>
```