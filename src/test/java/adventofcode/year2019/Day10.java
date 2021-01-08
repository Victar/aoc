package adventofcode.year2019;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day10 extends BaseTest {

	@Test public void singleCheck() {
		System.out.println(gcdByEuclidsAlgorithm(8, 1));
		System.out.println(gcdByEuclidsAlgorithm(6, 12));
		System.out.println(gcdByEuclidsAlgorithm(8, 0));
		System.out.println(gcdByEuclidsAlgorithm(10, 15));

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2019/day10/input.txt");
		final List<Asteroid> listAsteroid = parseList(data);//new ArrayList();
		for (String input : data) {
			System.out.println(input);
		}
		int max = Integer.MIN_VALUE;
//		int currentMax = availableCount(listAsteroid.get(3), listAsteroid);
		for (int i = 0; i < listAsteroid.size(); i++){
			int currentMax = availableCount(listAsteroid.get(i), listAsteroid);
			System.out.println(currentMax);
			max =Math.max(currentMax, max);
		}
		System.out.println("############");
		System.out.println(max);
	}

	public int availableCount(Asteroid current, final List<Asteroid> listAsteroid) {
		int count = 0;
		for (Asteroid asteroid : listAsteroid) {
			if (isVisible(current, asteroid, listAsteroid)) {
				count++;
			}
		}
		return count;
	}

	public boolean isVisible(Asteroid current, Asteroid destination, final List<Asteroid> listAsteroid) {
//		System.out.println("current - > destination "+ current +" -> " + destination);
		if (current.equals(destination)) {
			return false;
		}
		int xRange = current.getX() - destination.getX();
		int yRange = current.getY() - destination.getY();
		int gcd = gcdByEuclidsAlgorithm(xRange, yRange);
//		System.out.println(current + " " + destination + " xRange : " + xRange + " yRange: " + yRange + " gcd: " + gcd);
		boolean isVisible = true;
		if (Math.abs(gcd) >= 2) {
			//10, 15, 5
			int stepX = -xRange/gcd;
			int stepY = -yRange/gcd;
			if (gcd<0){
				 stepX = xRange/gcd;
				 stepY = yRange/gcd;
			}
			for (int i = 1; i < Math.abs(gcd); i++) {
				int checkX = current.getX() + stepX*i;
				int checkY = current.getY() + stepY*i;
//				System.out.println("checkX: " + checkX + " checkY: " +checkY);
				if (listAsteroid.contains(new Asteroid(checkX, checkY))){
					isVisible = false;
//					System.out.println("between: " +  new Asteroid(checkX, checkY));
				}
			}
		}
		return isVisible;
	}

	public boolean isVisible2(Asteroid current, Asteroid destination, final List<Asteroid> listAsteroid) {
//		System.out.println("current - > destination "+ current +" -> " + destination);
		if (current.equals(destination)) {
			return false;
		}
		boolean isVisible = true;
		for (Asteroid a : listAsteroid){
			if (collinear(current,destination,a)){
				isVisible = false;
			}
		}
		return isVisible;
	}


	public boolean collinear(Asteroid a1,Asteroid a2, Asteroid a3) {
	  return (a1.getY() - a2.getY()) * (a2.getX() - a3.getX()) == (a1.getY() - a3.getY()) * (a1.getX() - a2.getX());
	}

	public List<Asteroid> parseList(ArrayList<String> data) {
		final ArrayList<Asteroid> result = new ArrayList();
		for (int i = 0; i < data.size(); i++) {
			final String currentRow = data.get(i);
			char[] currentRowArra = currentRow.toCharArray();
			for (int j = 0; j < currentRowArra.length; j++) {
				if ('#' == currentRowArra[j]) {
					result.add(new Asteroid(i, j));
				}
			}
		}
		return result;
	}

	@Data
	class Asteroid {

		int x;
		int y;

		public Asteroid(final int x, final int y) {
			this.x = x;
			this.y = y;
		}

		@Override public String toString() {
			return "{" + "x=" + x + ", y=" + y + '}';
		}

		@Override public boolean equals(final Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;

			final Asteroid asteroid = (Asteroid) o;

			if (x != asteroid.x) return false;
			return y == asteroid.y;
		}

		@Override public int hashCode() {
			int result = x;
			result = 31 * result + y;
			return result;
		}
	}

	public int gcdByEuclidsAlgorithm(int n1, int n2) {
		if (n2 == 0) {
			return n1;
		}
		return gcdByEuclidsAlgorithm(n2, n1 % n2);
	}

}
