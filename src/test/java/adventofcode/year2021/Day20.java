package adventofcode.year2021;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.Set;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day20 extends BaseTest {

	int minX;
	int minY;
	int maxX;
	int maxY;
	String algString = StringUtils.EMPTY;
	boolean outsideLight;

	@Test public void runSilver() throws Exception {
		runAny(2);
	}

	@Test public void runGold() throws Exception {
		runAny(50);
	}

	public void runAny(final int rounds) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day20/input.txt");
		this.algString = data.get(0);
		this.maxX = data.get(2).length() - 1;
		this.maxY = data.size() - 3;
		final Set<Pixel> imageSet = new HashSet<>();

		for (int y = 2; y < data.size(); y++) {
			final String line = data.get(y);
			for (int x = 0; x < line.length(); x++) {
				if ("#".equals("" + line.charAt(x))) {
					final Pixel current = new Pixel(x, y - 2);
					imageSet.add(current);
				}
			}
		}
		Set<Pixel> result = imageSet;
		for (int i = 0; i < rounds; i++) {
			this.outsideLight = "#".equals("" + this.algString.charAt(0)) && i % 2 == 1;
			result = runTransform(result);
		}
		System.out.println(result.size());
		//		printImage(result);
	}

	public Set<Pixel> runTransform(final Set<Pixel> currentSet) {

		final Set<Pixel> result = new HashSet<>();
		for (int x = this.minX - 1; x <= this.maxX + 1; x++) {
			for (int y = this.minY - 1; y <= this.maxY + 1; y++) {
				final boolean isLight = isLight(currentSet, x, y);
				if (isLight) {
					result.add(new Pixel(x, y));
				}
			}
		}
		this.minX--;
		this.minY--;
		this.maxX++;
		this.maxY++;
		return result;
	}

	public void printImage(final Set<Pixel> imageSet) {
		for (int y = this.minY - 3; y <= this.maxY + 3; y++) {
			for (int x = this.minX - 3; x <= this.maxX + 3; x++) {
				final Pixel current = getPixel(imageSet, x, y);
				if (current == null) {
					System.out.print(" ");
				} else {
					System.out.print("#");
				}
			}
			System.out.println("   <-y" + y);
		}
	}

	public Pixel getPixel(final Set<Pixel> imageSet, final int finalX, final int finalY) {
		final Pixel toCheck = new Pixel(finalX, finalY);
		if (imageSet.contains(new Pixel(finalX, finalY))) {
			return toCheck;
		}
		return null;
	}

	public boolean isLight(final Set<Pixel> imageSet, final int currentX, final int currentY) {
		final StringBuilder sb = new StringBuilder();
		for (int y = currentY - 1; y <= currentY + 1; y++) {
			for (int x = currentX - 1; x <= currentX + 1; x++) {
				String cur = "0";
				if (getPixel(imageSet, x, y) != null) {
					cur = "1";
				} else {
					if (this.outsideLight && (x > this.maxX || x < this.minX || y > this.maxX || y < this.minY)) {
						cur = "1";
					}
				}
				sb.append(cur);
			}
		}
		return "#".equals("" + this.algString.charAt(Integer.parseInt(sb.toString(), 2)));
	}

	class Pixel {

		int x;
		int y;

		public Pixel(final int x, final int y) {
			this.x = x;
			this.y = y;
		}

		@Override public boolean equals(final Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;

			final Pixel pixel = (Pixel) o;

			if (this.x != pixel.x) return false;
			return this.y == pixel.y;
		}

		@Override public int hashCode() {
			int result = this.x;
			result = 31 * result + this.y;
			return result;
		}
	}

}
